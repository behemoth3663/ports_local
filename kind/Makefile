#

PORTNAME=	kind
PORTVERSION=	0.11.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@FreeBSD.org
COMMENT=	Kubernetes IN Docker - local clusters for testing Kubernetes

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
