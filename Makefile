GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
.PHONY: go-build
go-build: go-vendor
ifneq (${GIT_BRANCH}, main)
ifneq ($(shell git status --porcelain=v1 2>/dev/null | wc -l | tr -d " "), 0)
	go build -o spc -ldflags "-X github.com/dvdmuckle/spc/cmd.version=$(shell git rev-parse --short HEAD)-dirty"
else
	go build -o spc -ldflags "-X github.com/dvdmuckle/spc/cmd.version=$(shell git rev-parse --short HEAD)"
endif
else
	go build -o spc -ldflags "-X github.com/dvdmuckle/spc/cmd.version=$(shell git describe --tags --abbrev=0)"
endif
.PHONY: go-vendor
go-vendor:
	go mod vendor
clean:
	rm spc
	rm -rf rpm-build
	rm -rf rpm-build*
	rm -rf debbuild
.PHONY: rpm-build
rpm-build:
	mkdir -p rpm-build
ifdef mock-version
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=${GIT_BRANCH} --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build -r ${mock-version}
else 
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=${GIT_BRANCH} --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build
endif
.PHONY: rpm-build-all-arch
VERSION_ID = $(shell cat /etc/os-release | grep VERSION_ID | cut -d '=' -f2)
rpm-build-all-arch: $(shell ls /etc/mock/ | grep fedora-${VERSION_ID} | cut -d '-' -f3 | cut -d '.' -f1 | xargs -I ARCH echo -n rpm-build-all-arch.ARCH\ )
rpm-build-all-arch.%: ARCH=$*
rpm-build-all-arch.%:
	echo "Building RPM for ${ARCH}"
	$(MAKE) rpm-build mock-version=fedora-${VERSION_ID}-${ARCH} 
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=dev --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build-${ARCH} -r fedora-${VERSION_ID}-${ARCH}
	echo "Finished RPM for ${ARCH}"
rpm-build-docker:
# This has to run privileged as mock does some mounting stuff that doesn't work otherwise
	docker run --privileged -it -v $(CURDIR):/spc fedora /bin/bash -c "dnf install -y mock mock-scm make go-rpm-macros go-srpm-macros; cd spc; $(MAKE) rpm-build"
deb-build-docker:
	docker run --privileged -it -v $(CURDIR):/spc -v /run/user/1000/gnupg/S.gpg-agent:/run/user/1000/gnupg/S.gpg-agent -v ${HOME}/.gnupg:/root/.gnupg -e DEBIAN_FRONTEND=noninteractive -e GPG_TTY=`tty` ubuntu /bin/bash -c "ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime; apt-get update; apt-get install -y equivs devscripts dput make gpg software-properties-common; apt-add-repository -y ppa:longsleep/golang-backports; cd /spc; mk-build-deps --install -t 'apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y' debian/control; make prepare-deb-build"
prepare-deb-build: go-build
	./spc completion bash > debian/spc.bash-completion
	./spc docs --gen-tags=true  man spcdocs
	ls spcdocs > debian/manpages
	sed -i -e 's/^/spcdocs\//' debian/manpages
	$(MAKE) clean
	yes | debuild -S -d
	mkdir -p debbuild
	mv ../spc_* debbuild/
deb-bump-version:
	dch -i -M -D "noble"
bump-and-prepare-deb-build: deb-bump-version prepare-deb-build
rpm-bump-spec:
	rpmdev-bumpspec -u "David Muckle <dvdmuckle@dvdmuckle.xyz>" spc.spec
