# $FreeBSD$

PORTNAME=	awsls
PORTVERSION=	0.7.0
DISTVERSIONPREFIX=	v
CATEGORIES=	textproc

MAINTAINER=	none@freebsd.org
COMMENT=	List command for AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules
#GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	jckuester

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
