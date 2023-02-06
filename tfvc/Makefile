#

PORTNAME=	tfvc
PORTVERSION=	0.7.12
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform provider and module version checking tool
WWW=		https://tfverch.github.io/tfvc

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tfverch

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.version=${PORTVERSION}'"

fetch-version: .PHONY get-version

#post-patch: .PHONY .SILENT
#	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
