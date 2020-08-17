# $FreeBSD$

PORTNAME=	aws-ec2-instance-selector
PORTVERSION=	2.0.1
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils
DISTNAME=	amazon-ec2-instance-selector

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool which recommends instance types based on resource criteria

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_PROJECT=	${DISTNAME}

GO_TARGET=	./cmd/main.go:${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.sinclude "${.CURDIR}/Makefile.deps"
.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
