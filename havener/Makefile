#

PORTNAME=	havener
PORTVERSION=	2.0.8
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Swiss army knife for Kubernetes tasks
WWW=		https://github.com/homeport/havener

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static -X 'github.com/${GH_ACCOUNT}/${PORTNAME}/internal/cmd.version=${PORTVERSION}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#-		kubernetes-sigs:kustomize:v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
#+		kubernetes-sigs:kustomize:api/v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize \
#Not working well for unknown reason - when failed use patch- instead
GH_TUPLE:=	${GH_TUPLE:C|(kubernetes-sigs:kustomize:)(.*)/([^/]+)\$|\1\3/\2|}

.include <bsd.port.mk>
