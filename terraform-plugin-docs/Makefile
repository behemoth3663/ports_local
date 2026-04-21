PORTNAME=	terraform-plugin-docs
PORTVERSION=	0.25.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Generate and validate Terraform plugin/provider documentation
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

GO_TARGET=	./cmd/tfplugindocs

PLIST_FILES=	bin/tfplugindocs

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

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
