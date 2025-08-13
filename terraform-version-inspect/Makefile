PORTNAME=	terraform-version-inspect
PORTVERSION=	0.2.8
DISTVERSIONPREFIX=	v
#PORTREVISION=	19
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Determine required terraform version
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	ikorchynskyi

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w  -extldflags '-static'"

post-extract: .PHONY
	${TOUCH} ${WRKSRC}/LICENSE

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)|docker:docker-credential-helpers\1|}

.include <bsd.port.mk>
