# SPEC file

%global c_vendor    %{_vendor}
%global gh_owner    %{_owner}
%global gh_project  %{_project}

Name:      %{_package}
Version:   %{_version}
Release:   %{_release}%{?dist}
Summary:   NATS bus ping command.

Group:     Applications/Services
License:   Apache-2.0
URL:       https://github.com/%{gh_owner}/%{gh_project}

BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-%(%{__id_u} -n)

Provides:  %{gh_project} = %{version}

%description
Ping a NATS bus to see if it is alive.

%build
(cd %{_current_directory} && make build)

%install
rm -rf $RPM_BUILD_ROOT
(cd %{_current_directory} && make install DESTDIR=$RPM_BUILD_ROOT)

%clean
rm -rf $RPM_BUILD_ROOT
(cd %{_current_directory} && make clean)

%files
%attr(-,root,root) %{_binpath}
%attr(-,root,root) %{_docpath}
%attr(-,root,root) %{_manpath}
%docdir %{_docpath}
%docdir %{_manpath}
%config(noreplace) %{_configpath}*

%changelog
* Thu Jul 28 2016 Nicola Asuni <nicola.asuni@miracl.com> 1.0.0-1
- Initial Commit
