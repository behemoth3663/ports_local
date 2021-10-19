PORTNAME=	hsbench
PORTVERSION=	0.2+
CATEGORIES=	net

MAINTAINER=	none@FreeBSD.org
COMMENT=	S3 benchmark utility

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	vitalif
GH_TAGNAME=	8782423

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
