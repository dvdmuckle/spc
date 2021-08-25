.PHONY: go-build
go-build: go-vendor
	go build -o spc
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
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=dev --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build -r ${mock-version}
else 
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=dev --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build
endif
.PHONY: rpm-build-all-arch
VERSION_ID = $(shell cat /etc/os-release | grep VERSION_ID | cut -d '=' -f2)
rpm-build-all-arch: $(shell ls /etc/mock/ | grep fedora-${VERSION_ID} | cut -d '-' -f3 | cut -d '.' -f1 | xargs -I ARCH echo -n rpm-build-all-arch.ARCH\ )
rpm-build-all-arch.%: ARCH=$*
rpm-build-all-arch.%:
	echo "Building RPM for ${ARCH}"
	mock --scm-enable --scm-option method=git --scm-option package=spc --scm-option spec=spc.spec --scm-option branch=dev --scm-option write_tar=True --scm-option git_get='git clone https://github.com/dvdmuckle/spc.git' --enable-network --resultdir rpm-build-${ARCH} -r fedora-${VERSION_ID}-${ARCH}
	echo "Finished RPM for ${ARCH}"
rpm-build-docker:
# This has to run privileged as mock does some mounting stuff that doesn't work otherwise
	docker run --privileged -it -v $(CURDIR):/spc fedora /bin/bash -c "dnf install -y mock mock-scm make go-rpm-macros go-srpm-macros; cd spc; $(MAKE) rpm-build"
prepare-deb-build: go-build
	./spc completion bash > debian/spc.bash-completion
	$(MAKE) clean
	cd debian
	cd ..
	yes | debuild -S -d
	mkdir -p debbuild
	mv ../spc_* debbuild/
deb-bump-version:
	dch -i -M -D focal
bump-and-prepare-deb-build: deb-bump-version prepare-deb-build
