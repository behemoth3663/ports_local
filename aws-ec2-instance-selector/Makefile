#

PORTNAME=	aws-ec2-instance-selector
PORTVERSION=	2.4.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
DISTNAME=	amazon-ec2-instance-selector

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool which recommends instance types based on resource criteria
WWW=		https://github.com/aws/amazon-ec2-instance-selector

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_PROJECT=	${DISTNAME}

GO_TARGET=	./cmd/main.go:${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
