PORTNAME=	tgcon
PORTVERSION=	20221121
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Generates struct field tag values as string constants
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	amarjeetanandsingh
GH_TAGNAME=	6b3f725

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v3/pkg/version.BuildVersion=${DISTVERSIONPREFIX}${PORTVERSION}'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
