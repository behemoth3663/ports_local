#

PORTNAME=	hsbench
PORTVERSION=	2.0.5
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	none@FreeBSD.org
COMMENT=	S3 benchmark utility

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	amilevskiy

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}'"

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
