PORTNAME=	havener
PORTVERSION=	2.2.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Swiss army knife for Kubernetes tasks
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules,1.23

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal/cmd.version=${PORTVERSION}'"

pre-deps: .PHONY .SILENT
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# Commented out due to using hash (not suffix/version) that is working
#-		kubernetes-sigs:kustomize:v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#+		kubernetes-sigs:kustomize:api/v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#GH_TUPLE:=	${GH_TUPLE:C|^(kubernetes-sigs:kustomize:)([^:]+)(:.*/vendor/sigs\\.k8s\\.io/kustomize)/([^[:space:]]+)|\1\4/\2\3|}

.include <bsd.port.mk>
