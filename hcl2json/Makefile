PORTNAME=	hcl2json
PORTVERSION=	0.6.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Convert hcl2 to json
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tmccombs

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${PORTVERSION}'"

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
