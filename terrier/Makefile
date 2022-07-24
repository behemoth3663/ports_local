#

PORTNAME=	terrier
PORTVERSION=	0.0.2
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Tool for Identifying and Analyzing Container and Image Components

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	heroku

EXAMPLESDIR=	${PREFIX}/share/examples/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME} \
		share/examples/${PORTNAME}/cfg.yml

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${PREFIX}/share/examples/${PORTNAME}
	${INSTALL_DATA} ${WRKSRC}/cfg.yml ${STAGEDIR}${PREFIX}/share/examples/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
