# $FreeBSD$

PORTNAME=	phoneinfoga
PORTVERSION=	2.3.7
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Advanced information gathering & OSINT framework for phone numbers

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sundowndev
GH_PROJECT=	PhoneInfoga

PLIST_FILES=	bin/${PORTNAME}

#doesn't work
#GO_BUILDFLAGS+=	-ldflags="-X config.Version=${PORTVERSION}"

post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's/^[[:space:]]*(var[[:space:]]+Version[[:space:]]*=).*/\1 "${DISTVERSIONPREFIX}${PORTVERSION}"/' \
		-e 's/^[[:space:]]*(var[[:space:]]+Commit[[:space:]]*=).*/\1 "f52c5338fd9d6711e5375e53a5356890680b31c7"/' \
	${WRKSRC}/config/version.go

pre-build: .PHONY
	${MKDIR} ${WRKSRC}/client/dist
	cd ${WRKSRC}/api && \
	${GO_CMD} run github.com/jessevdk/go-assets-builder ../client/dist -o ./assets.go -p api && \
	${GO_CMD} run ${WRKSRC}/vendor/github.com/swaggo/swag/cmd/swag/main.go init -g ./server.go --parseDependency

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
