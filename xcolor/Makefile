PORTNAME=	xcolor
DISTVERSION=	0.5.1
#PORTREVISION=	1
CATEGORIES=	local x11

MAINTAINER=	none@freebsd.org
COMMENT=	Lightweight color picker for X11
WWW=		https://soft.github.io/xcolor

LICENSE=	MIT

USES=		cargo xorg desktop-file-utils
USE_XORG=	xcb

USE_GITHUB=	yes
GH_ACCOUNT=	Soft

PLIST_FILES= \
	bin/${PORTNAME} \
	share/man/man1/${PORTNAME}.1.gz \
	share/applications/${PORTNAME}.desktop \
	share/icons/hicolor/16x16/apps/${PORTNAME}.png \
	share/icons/hicolor/24x24/apps/${PORTNAME}.png \
	share/icons/hicolor/32x32/apps/${PORTNAME}.png \
	share/icons/hicolor/48x48/apps/${PORTNAME}.png \
	share/icons/hicolor/256x256/apps/${PORTNAME}.png \
	share/icons/hicolor/512x512/apps/${PORTNAME}.png

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
	${INSTALL_MAN} ${WRKSRC}/man/${PORTNAME}.1 ${STAGEDIR}${PREFIX}/share/man/man1
	${INSTALL_DATA} ${WRKSRC}/extra/${PORTNAME}.desktop ${STAGEDIR}${PREFIX}/share/applications
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/16x16/apps
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/24x24/apps
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/32x32/apps
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/48x48/apps
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/256x256/apps
	${MKDIR} ${STAGEDIR}${PREFIX}/share/icons/hicolor/512x512/apps
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-16.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/16x16/apps/${PORTNAME}.png
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-24.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/24x24/apps/${PORTNAME}.png
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-32.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/32x32/apps/${PORTNAME}.png
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-48.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/48x48/apps/${PORTNAME}.png
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-256.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/256x256/apps/${PORTNAME}.png
	${INSTALL_DATA} ${WRKSRC}/extra/icons/${PORTNAME}-512.png ${STAGEDIR}${PREFIX}/share/icons/hicolor/512x512/apps/${PORTNAME}.png

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > Makefile.crates )
	( ${MAKE} -C ${.CURDIR} makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
