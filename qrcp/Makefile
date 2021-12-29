#

PORTNAME=	qrcp
PORTVERSION=	0.8.4
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Transfer files over Wi-Fi from computer to mobile device

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	claudiodangelis

#didn't work
#GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
#		-X version.version=${PORTVERSION} \
#		-X version.date=${%FT%TZ:L:gmtime}'"

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*version[[:space:]]*=[[:space:]]*).*|\1"${PORTVERSION}"|' \
		-e 's|^([[:space:]]*date[[:space:]]*=[[:space:]]*).*|\1"${%FT%TZ:L:gmtime}"|' \
	${WRKSRC}/version/version.go

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor &&  make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
