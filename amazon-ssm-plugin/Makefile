PORTNAME=	amazon-ssm-plugin
PORTVERSION=	1.2.764.0
PORTREVISION=	3
CATEGORIES=	local sysutils

MAINTAINER=	nork@FreeBSD.org
COMMENT=	Amazon AWS Manages shell experience using SSM APIs
WWW=		https://github.com/${GH_ACCOUNT}/session-manager-plugin

LICENSE=	APACHE20

RUN_DEPENDS=	aws:devel/py-awscli@${PY_FLAVOR}

USES=		go:modules python:env

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"
GO_TARGET=	src/sessionmanagerplugin-main/main.go:${PREFIX}/bin/session-manager-plugin \
		src/ssmcli-main/main.go:${PREFIX}/bin/ssmcli

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_PROJECT=	session-manager-plugin
GH_SUBDIR=	${GH_PROJECT}

PORTDOCS=	NOTICE README.md RELEASENOTES.md THIRD-PARTY VERSION
PLIST_FILES=	${ETCDIR}/sessionmanagerplugin/seelog.xml.template \
		${ETCDIR}/SSMCLI/seelog.xml.template \
		bin/session-manager-plugin \
		bin/ssmcli \
		'@dir /var/log/amazon-ssm-plugin'

LOGDIR=		${DESTDIR}/var/log/${PORTNAME}

OPTIONS_DEFINE=	DOCS

post-patch: .PHONY .SILENT
	${CP} -a ${PATCHDIR}/go.mod ${PATCHDIR}/go.sum ${WRKSRC}
	${REINPLACE_CMD} -e 's|%%VERSION%%|${DISTVERSION}|' ${WRKSRC}/src/version/version.go
	${REINPLACE_CMD} -e 's|%%ETCDIR%%|${ETCDIR}|g; s|%%LOGDIR%%|${LOGDIR}|g' ${WRKSRC}/seelog_unix.xml
	${REINPLACE_CMD} -e 's|%%ETCDIR%%|${ETCDIR}|g; s|%%LOGDIR%%|${LOGDIR}|g' ${WRKSRC}/src/log/log_unix.go

post-install: .PHONY
	@${MKDIR} ${STAGEDIR}${LOGDIR} ${STAGEDIR}${ETCDIR}/SSMCLI ${STAGEDIR}${ETCDIR}/sessionmanagerplugin
	${INSTALL_DATA} ${WRKSRC}/seelog_unix.xml ${STAGEDIR}${ETCDIR}/SSMCLI/seelog.xml.template
	${INSTALL_DATA} ${WRKSRC}/seelog_unix.xml ${STAGEDIR}${ETCDIR}/sessionmanagerplugin/seelog.xml.template

post-install-DOCS-on: .PHONY
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	cd ${WRKSRC} && ${INSTALL_DATA} ${PORTDOCS} ${STAGEDIR}${DOCSDIR}

pre-deps: .PHONY .SILENT
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v
	${TAR} --create --directory ${PATCHDIR} --exclude='go.mod' --exclude='go.sum' --xz --file ${TMPDIR}/${PORTNAME}.tar.xz .
	${RM} ${PATCHDIR}/*
	${CP} -a ${WRKSRC}/go.mod ${WRKSRC}/go.sum ${PATCHDIR}

pre-deps-all: .PHONY
	${TAR} --extract --directory ${PATCHDIR} --file ${TMPDIR}/${PORTNAME}.tar.xz
	${RM} ${PATCH_COOKIE} ${TMPDIR}/${PORTNAME}.tar.xz
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

post-deps-all: .PHONY

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-twinj:uuid:v1.0.0:twinj_uuid
GH_TUPLE:=	${GH_TUPLE:Ntwinj\:uuid\:*}

.include <bsd.port.mk>
