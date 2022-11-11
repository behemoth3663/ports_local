# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	kubectl
#PORTVERSION=	1.20.11, 1.22.2, 1.21.5
PORTVERSION=	1.25.4
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes command-line tool
WWW=		https://github.com/kubernetes/kubectl

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		cpe go:modules
CPE_VENDOR=	jenkins
CPE_PRODUCT=	kubernetes

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes
GH_PROJECT=	kubernetes

GO_PKGNAME=	k8s.io/kubernetes
GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${GH_PROJECT}-${PORTVERSION}^{}

GO_BUILDFLAGS=	-ldflags="-s -w -extldflags '-static' \
		-X k8s.io/component-base/version.buildDate=${%FT%TZ:L:gmtime} \
		-X k8s.io/component-base/version.gitTreeState= \
		-X k8s.io/component-base/version.gitCommit=${:!${GH_COMMIT}!:[1]} \
		-X k8s.io/component-base/version.gitMajor=${PORTVERSION:R:R} \
		-X k8s.io/component-base/version.gitMinor=${PORTVERSION:R:E} \
		-X k8s.io/component-base/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/client-go/pkg/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

#to get the new version value
#KUBERNETES_API_VERSION!=git ls-remote --tags \
#		https://github.com/kubernetes/api | \
#		sed -rn 's|^.*[[:space:]]refs/tags/(${DISTVERSIONPREFIX}[.0-9]+)\^\{\}[[:space:]]*$$|\1|p' | \
#		sort --version-sort --reverse | head -1
KUBERNETES_API_VERSION=	v0.25.4

OPTIONS_DEFINE=		BASH ZSH
OPTIONS_DEFAULT=	BASH ZSH

BASH_PLIST_FILES=	share/bash-completion/completions/${PORTNAME}
ZSH_PLIST_FILES=	share/zsh/site-functions/_${PORTNAME}

post-install-BASH-on: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/bash-completion/completions
	${STAGEDIR}${PREFIX}/bin/${PORTNAME} completion bash > ${STAGEDIR}${PREFIX}/share/bash-completion/completions/${PORTNAME}

post-install-ZSH-on: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/zsh/site-functions
	${STAGEDIR}${PREFIX}/bin/${PORTNAME} completion zsh > ${STAGEDIR}${PREFIX}/share/zsh/site-functions/_${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

_OPTIONS_extract+=	600:pre-post-extract
pre-post-extract: .PHONY
	@test -z '${WRKSRC_docopt_docopt_go}' || \
		${MV} ${WRKSRC_docopt_docopt_go:S/docopt-go/docopt.go/} ${WRKSRC_docopt_docopt_go}

.include <bsd.port.mk>
