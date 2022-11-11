PORTNAME=	helm
PORTVERSION=	3.10.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes Package Manager
WWW=		https://helm.sh

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git
RUN_DEPENDS=	kubectl:sysutils/kubectl

USES=		cpe go:modules tar:xz

USE_GITHUB=	yes
GH_ACCOUNT=	${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_PKGNAME=	k8s.io/helm
GO_BUILDFLAGS=	-ldflags="-s -w -extldflags '-static' \
		-X helm.sh/helm/v3/internal/version.version=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X helm.sh/helm/v3/internal/version.gitCommit=${:!${GH_COMMIT}!:[1]} \
		-X helm.sh/helm/v3/internal/version.gitTreeState=clean"

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

OPTIONS_DEFINE=		BASH FISH ZSH
OPTIONS_DEFAULT=	BASH FISH ZSH

BASH_PLIST_FILES=	share/bash-completion/completions/helm
FISH_PLIST_FILES=	share/fish/vendor_functions.d/helm.fish
ZSH_PLIST_FILES=	share/zsh/site-functions/_helm

post-install-BASH-on:
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/bash-completion/completions
	${STAGEDIR}${PREFIX}/bin/helm completion bash > ${STAGEDIR}${PREFIX}/share/bash-completion/completions/helm

post-install-FISH-on:
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/fish/vendor_functions.d
	${STAGEDIR}${PREFIX}/bin/helm completion fish > ${STAGEDIR}${PREFIX}/share/fish/vendor_functions.d/helm.fish

post-install-ZSH-on:
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/zsh/site-functions
	${STAGEDIR}${PREFIX}/bin/helm completion zsh > ${STAGEDIR}${PREFIX}/share/zsh/site-functions/_helm

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#-kubernetes-sigs:kustomize:v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#-kubernetes-sigs:kustomize:v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml
#+kubernetes-sigs:kustomize:api/v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize
#+kubernetes-sigs:kustomize:kyaml/v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize
GH_TUPLE:=	${GH_TUPLE:C|^(kubernetes-sigs:kustomize:)([^:]+)(:.*/vendor/sigs\\.k8s\\.io/kustomize)/([^[:space:]]+)|\1\4/\2\3|}

#-moby:moby:v0.6.4:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.6.4:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)|docker:docker-credential-helpers\1|}

#-etcd-io:etcd:v3.5.4:etcd_io_etcd/vendor/go.etcd.io/etcd/api/v3
#+etcd-io:etcd:v3.5.4:etcd_io_etcd/vendor/go.etcd.io/etcd
GH_TUPLE:=	${GH_TUPLE:C|^(etcd-io:etcd:.*)/api/v3|\1|}

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@test ! -e ${WRKSRC_etcd_io_etcd}/api || \
		${RLN} ${WRKSRC_etcd_io_etcd}/api ${WRKSRC}/vendor/go.etcd.io/etcd/api/v3

.include <bsd.port.mk>
