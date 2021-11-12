#

PORTNAME=	infracost
PORTVERSION=	0.9.13
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Cloud cost estimates for Terraform

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	infracost

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

pre-build: .PHONY
	cd ${WRKSRC} && ${GO_CMD} mod vendor -v && ${GO_CMD} mod tidy -v

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
#.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
