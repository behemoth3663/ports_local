PORTNAME=	pike
PORTVERSION=	0.3.24
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Determine permissions or policy required for IAC code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jamesWoolfenden

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^(const Version).*$$|\1 = "${PORTVERSION}"|' \
	${WRKSRC}/src/version.go

pre-deps pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
