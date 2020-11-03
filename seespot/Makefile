# Created by: a1vm
# $FreeBSD: head/local/seespot/Makefile 000000 2018-08-22 06:54:28Z meta $

PORTNAME=	seespot
PORTVERSION=	0.0.1
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Watches for a termination notification in AWS

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go
USE_GITHUB=	yes
GH_ACCOUNT=	opszero

#No longer maintained
GH_TAGNAME=	8280e32

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${PORTNAME}
PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
