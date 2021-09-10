# Generated by go2rpm 1
%bcond_without check

# https://github.com/dvdmuckle/spc

%global goipath         github.com/dvdmuckle/spc
%global tag             0.10.2
Version:                %{tag}
%gometa


%global common_description %{expand:
A lightweight multiplatform CLI for Spotify.}

%global godocs          README.md

Name:           spc
Release:        1%{?dist}
Summary:        A lightweight multiplatform CLI for Spotify

# Upstream license specification: Apache-2.0
License:        ASL 2.0
URL:            %{gourl}
Source0:        %{gosource}

# Using go mod vendor to get the build requirements, since
# we would have to anyways for all the packages that don't
# have Fedora packages
# BuildRequires:  golang(github.com/golang/glog)
# Package does not exist: BuildRequires:  golang(github.com/ktr0731/go-fuzzyfinder)
# Package does not exist: BuildRequires:  golang(github.com/markbates/goth/providers/spotify)
# BuildRequires:  golang(github.com/mitchellh/go-homedir)
# BuildRequires:  golang(github.com/spf13/cobra)
# BuildRequires:  golang(github.com/spf13/viper)
# Package does not exist: BuildRequires:  golang(github.com/zmb3/spotify)
# BuildRequires:  golang(golang.org/x/oauth2)
BuildRequires:  git
BuildRequires:  dbus-X11

Requires: bash-completion

%description
%{common_description}

%gopkg

%prep
%goprep

%build
go mod vendor
export LDFLAGS="-X %{goipath}/cmd.version=%{tag}"
%gobuild -o %{gobuilddir}/bin/spc %{goipath}
%{gobuilddir}/bin/spc completion bash > %{gobuilddir}/spc.bash
%{gobuilddir}/bin/spc completion zsh > %{gobuilddir}/spc.zsh
%{gobuilddir}/bin/spc docs --gen-tags=true man %{gobuilddir}/spcdocs


%install
%gopkginstall
install -m 0755 -vd                     %{buildroot}%{_bindir}
install -m 0755 -vp %{gobuilddir}/bin/* %{buildroot}%{_bindir}/
mkdir -p %{buildroot}/usr/share/bash-completion/completions
mkdir -p %{buildroot}/usr/share/zsh/site-functions
mkdir -p %{buildroot}%{_mandir}/man1
install -m 0644 -vp %{gobuilddir}/spc.bash %{buildroot}/usr/share/bash-completion/completions/spc
install -m 0644 -vp %{gobuilddir}/spc.zsh %{buildroot}/usr/share/zsh/site-functions/_spc
install -m 0644 -vpt %{buildroot}%{_mandir}/man1/ %{gobuilddir}/spcdocs/spc* 



%if %{with check}
%check
%gocheck
%endif

%files
%doc README.md
%license LICENSE
%{_bindir}/*
/usr/share/bash-completion/completions/spc
/usr/share/zsh/site-functions/_spc
%{_mandir}/man1/spc*


%gopkgfiles

%changelog
* Mon Sep 06 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.10.2-1
- Specify device on all playback commands
- Check owner of Discover Weekly playlist before saving

* Sun Sep 05 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.10.1-1
- Seek command now accepts timestamps

* Thu Sep 02 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.10.0-1
- Man pages and docs command

* Sun Aug 29 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.9.0-1
- Add version command

* Thu Aug 26 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.8.1-1
- More clarity on how to treat auth link for terminals that don't support links

* Thu Aug 26 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.8.0-1
- New save-weekly command
- Change in-memory token refresh to only refresh if token expired

* Tue Aug 24 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.7.1-1
- Fix typo in search command help

* Thu Aug 19 2021 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.7.0-2
- Allows for searching for artist

* Mon Aug 24 14:01:00 EDT 2020 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.5.1-1
- Fix pathing for config file in SetupConfig()

* Sun Aug 23 16:46:00 EDT 2020 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.5-1
- Add config subcommand

* Sat Aug 22 13:36:00 EDT 2020 David Muckle <dvdmuckle@dvdmuckle.xyz> - 0.4-1
- Initial package