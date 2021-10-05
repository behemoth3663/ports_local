# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	kubectl
#PORTVERSION=	1.20.11, 1.22.2
PORTVERSION=	1.21.5
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes command-line tool

LICENSE=	APACHE20

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes
GH_PROJECT=	kubernetes

GH_HASH=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

GO_PKGNAME=	k8s.io/kubernetes
GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}

GO_BUILDFLAGS=	-ldflags="-s -w -buildid= \
		-X k8s.io/component-base/version.buildDate=${:!date +%Y-%m-%dT%H:%M:%SZ!} \
		-X k8s.io/component-base/version.gitTreeState=clean \
		-X k8s.io/component-base/version.gitCommit=${:!${GH_HASH}!:[1]} \
		-X k8s.io/component-base/version.gitMajor=${PORTVERSION:R:R} \
		-X k8s.io/component-base/version.gitMinor=${PORTVERSION:R:E} \
		-X k8s.io/component-base/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/client-go/pkg/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION}"
#GO_BUILDFLAGS+=	-installsuffix static
#GO_BUILDFLAGS+=	-asmflags -trimpath=${WRKSRC}
#GO_BUILDFLAGS+=	-gcflags -trimpath=${WRKSRC}

PLIST_FILES=	bin/${PORTNAME}

# The vendored golang/sys fails to build on aarch64
#GOLANG_SYS_V=	22da62e12c0c
#GH_TUPLE=	golang:sys:${GOLANG_SYS_V}:golang_sys/vendor/golang.org/x/sys_${GOLANG_SYS_V}
#_OPTIONS_extract+=	750:post-post-extract
#post-post-extract: .PHONY
#	${RM} -r ${WRKSRC}/vendor/golang.org/x/sys
#	${MV} ${WRKSRC}/vendor/golang.org/x/sys_${GOLANG_SYS_V} ${WRKSRC}/vendor/golang.org/x/sys

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

.include <bsd.port.pre.mk>
#GO_CMD=	${LOCALBASE}/go-1.16/bin/go
GO_BUILDFLAGS:=	${GO_BUILDFLAGS:N-mod=vendor} -mod=mod
.include <bsd.port.post.mk>
