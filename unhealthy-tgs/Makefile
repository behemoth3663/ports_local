PORTNAME=	unhealthy-tgs
PORTVERSION=	0.1.13
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Get unhealthy target groups of Load Balancers
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}/tree/main/cmd/${PORTNAME}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi
GH_PROJECT=	go

WRKSRC_SUBDIR=	cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
		-X main.version=${PORTVERSION} \
		-X main.commit=$$(${COMMIT_CMD})"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY .SILENT
	${RM} ${WRKSRC}/../../go.work
	${MKDIR} ${WRKSRC}/vendor
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

post-patch: .PHONY .SILENT
	test ! -f ${PATCHDIR}/go-spew.patch || ${PATCH} --directory ${WRKSRC}/../.. --strip 1 < ${PATCHDIR}/go-spew.patch

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}
GH_TUPLE:=	${GH_TUPLE:Ndavecgh\:go-spew\:*}

.include <bsd.port.mk>
