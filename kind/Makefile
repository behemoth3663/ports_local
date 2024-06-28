PORTNAME=	kind
PORTVERSION=	0.23.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@FreeBSD.org
COMMENT=	Kubernetes IN Docker - local clusters for testing Kubernetes
WWW=		https://kind.sigs.k8s.io/

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
