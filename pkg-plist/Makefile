# Created by: Martin Tournoij <martin@arp242.net>
# $FreeBSD: head/ports-mgmt/pkg-plist/Makefile 512125 2019-09-15 17:27:41Z antoine $

PORTNAME=	pkg-plist
PORTVERSION=	1.4
PORTEPOCH=	9
CATEGORIES=	local ports-mgmt python
MASTER_SITES=	http://code.arp242.net/${BB_PROJECT}/get/${BB_COMMIT}.tar.gz?dummy=/

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Generate a pkg-plist for a port

LICENSE=	MIT

USES=		python:run shebangfix
SHEBANG_FILES=	pkg-plist
NO_BUILD=	yes
WRKSRC=		${WRKDIR}/${BB_ACCOUNT}-${BB_PROJECT}-${BB_COMMIT}

BB_ACCOUNT=	Carpetsmoker
BB_PROJECT=	${PORTNAME}
BB_COMMIT=	ff040a64c20b

NO_ARCH=	yes

PLIST_FILES=	bin/pkg-plist share/man/man1/pkg-plist.1.gz

do-install:
	${INSTALL_SCRIPT} ${WRKSRC}/pkg-plist ${STAGEDIR}${PREFIX}/bin
	${INSTALL_MAN} ${WRKSRC}/pkg-plist.1 ${STAGEDIR}${PREFIX}/share/man/man1

.include <bsd.port.mk>
