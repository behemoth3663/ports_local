#

PORTNAME=	traefik-config-validator
PORTVERSION=	0.0.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Validate your Traefik configuration files

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	otto-de

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'github.com/main.version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

post-patch:
	${REINPLACE_CMD} -r \
		's|^([[:space:]]*version[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
	${WRKSRC}/${GO_TARGET}/main.go

.include <bsd.port.mk>
