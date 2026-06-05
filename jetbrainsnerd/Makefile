PORTNAME=	jetbrainsnerd
PORTVERSION=	3.4.0
#PORTEPOCH=	1
CATEGORIES=	local x11-fonts
MASTER_SITES=	https://github.com/ryanoasis/nerd-fonts/releases/download/v${PORTVERSION}/
DISTNAME=	JetBrainsMono

MAINTAINER=	none@freebsd.org
COMMENT=	Free and open source typeface for developers with glyphs
WWW=		https://www.nerdfonts.com/

LICENSE=	OFL11
LICENSE_FILE=	${WRKSRC}/OFL.txt

USES=		fonts zip

NO_ARCH=	yes
NO_BUILD=	yes
NO_WRKSUBDIR=	yes

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${FONTSDIR}
	${INSTALL_DATA} ${WRKSRC}/*.ttf ${STAGEDIR}${FONTSDIR}
	${LS} ${STAGEDIR}${FONTSDIR}/*.ttf | ${SED} -e's,^${STAGEDIR},,' >> ${TMPPLIST}

.include <bsd.port.mk>
