PORTNAME=	ytt
PORTVERSION=	0.51.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML tool with templates and user provided values
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	k14s

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w  -extldflags '-static' \
	-X 'github.com/vmware-tanzu/carvel-ytt/pkg/version.Version=${PORTVERSION}'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
