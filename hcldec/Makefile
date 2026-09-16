PORTNAME=	hcldec
PORTVERSION=	2.25.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Transforms HCL into JSON using decoding specification
WWW=		https://github.com/${GH_ACCOUNT}/hcl

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp
GH_PROJECT=	hcl

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

#const versionStr = "0.0.1-dev"
post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${REINPLACE_CMD} -r \
		-e 's|^([[:blank:]]*const[[:blank:]]+versionStr[[:blank:]]*=[[:blank:]]*)".*"[[:blank:]]*$$|\1"${PORTVERSION}"|' \
	${WRKSRC}/${GO_TARGET}/main.go

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
