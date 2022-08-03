#

PORTNAME=	terraform-module-versions
PORTVERSION=	3.1.13
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool that checks Terraform code for module updates

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	keilerkonzept

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.version=${PORTVERSION}'"

fetch-version: .PHONY get-version

post-patch: .PHONY .SILENT
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
