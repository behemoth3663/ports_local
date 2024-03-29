# Created by: Toni Viemer
# $FreeBSD: head/security/makepasswd/Makefile 530719 2020-04-04 20:51:40Z antoine $

PORTNAME=	makepasswd
PORTVERSION=	1.10
PORTREVISION=	6
CATEGORIES=	local security
MASTER_SITES=	DEBIAN_POOL
DISTNAME=	${PORTNAME}_${PORTVERSION}.orig

PATCHFILES=	${PORTNAME}_${PORTVERSION}-5.diff.gz
PATCH_SITES=	${MASTER_SITES}
PATCH_DIST_STRIP=	-p1

MAINTAINER=	toni.viemero@iki.fi
COMMENT=	Random password generator

DEPRECATED=	Broken for more than 6 months
EXPIRATION_DATE=	2020-05-05

LICENSE=	GPLv1

NO_BUILD=	yes
USES=		perl5 shebangfix
SHEBANG_FILES=	makepasswd
USE_PERL5=	run

OPTIONS_DEFINE=	DOCS

RUN_DEPENDS+=	p5-Crypt-PasswdMD5>=0:security/p5-Crypt-PasswdMD5 \
		p5-Crypt-OpenSSL-Random>=0:security/p5-Crypt-OpenSSL-Random

WRKSRC=		${WRKDIR}/${PORTNAME}-${PORTVERSION}

PLIST_FILES=	bin/${PORTNAME} man/man1/${PORTNAME}.1.gz
PORTDOCS=		CHANGES README

post-patch:
	${REINPLACE_CMD} -e "s/use	bytes;/#use	bytes;/" ${WRKSRC}/makepasswd

do-install:
	${INSTALL_SCRIPT} ${WRKSRC}/makepasswd ${STAGEDIR}${PREFIX}/bin
	${INSTALL_MAN} ${WRKSRC}/${PORTNAME}.1 ${STAGEDIR}${PREFIX}/man/man1
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
.for file in ${PORTDOCS}
	${INSTALL_DATA} ${WRKSRC}/${file} ${STAGEDIR}${DOCSDIR}
.endfor

.include <bsd.port.mk>
