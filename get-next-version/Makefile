PORTNAME=	get-next-version
PORTVERSION=	2.7.3
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Gets next semantic version based on conventional commits
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	thenativeweb

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONFULL}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${DISTVERSIONFULL}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitVersion=${GIT_COMMIT:sh:C/(.{7}).*/\\1/W}'"

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
