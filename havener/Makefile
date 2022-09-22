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

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#-		kubernetes-sigs:kustomize:v0.11.0:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#-		kubernetes-sigs:kustomize:v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
#+		kubernetes-sigs:kustomize:kyaml/v0.11.0:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#+		kubernetes-sigs:kustomize:api/v0.8.11:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
GH_TUPLE:=	${GH_TUPLE:C|(kubernetes-sigs:kustomize:)(.*)/([^/]+)\$|\1\3/\2/\3|}

.include <bsd.port.mk>
