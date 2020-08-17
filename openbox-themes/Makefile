# Created by: Roman Bogorodskiy <novel@FreeBSD.org>
# $FreeBSD: head/x11-themes/openbox-themes/Makefile 530719 2020-04-04 20:51:40Z antoine $

PORTNAME=	openbox-themes
PORTVERSION=	1.0.2
CATEGORIES=	local x11-themes
MASTER_SITES=	DEBIAN
DISTNAME=	${PORTNAME}_${PORTVERSION}

MAINTAINER=	novel@FreeBSD.org
COMMENT=	Themes for Openbox

DEPRECATED=	Broken for more than 6 months
EXPIRATION_DATE=	2020-05-05

LICENSE=	GPLv2

RUN_DEPENDS=	openbox:x11-wm/openbox

NO_BUILD=	yes
NO_ARCH=	yes
WRKSRC=		${WRKDIR}/${PORTNAME}-${PORTVERSION}

post-extract:
	@${RM} -r ${WRKSRC}/debian

do-install:
	@${MKDIR} ${STAGEDIR}${PREFIX}/share/themes
	@cd ${WRKSRC} && ${COPYTREE_SHARE} . ${STAGEDIR}${PREFIX}/share/themes

.include <bsd.port.mk>
