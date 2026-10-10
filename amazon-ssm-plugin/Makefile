PORTNAME=	amazon-ssm-plugin
PORTVERSION=	1.2.896.0
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	nork@FreeBSD.org
COMMENT=	Amazon AWS Manages shell experience using SSM APIs
WWW=		https://github.com/${GH_ACCOUNT}/session-manager-plugin

LICENSE=	APACHE20

RUN_DEPENDS=	aws:devel/py-awscli@${PY_FLAVOR}

USES=		go:modules python:env

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_PROJECT=	session-manager-plugin
GH_SUBDIR=	${GH_PROJECT}

GO_TARGET=	src/sessionmanagerplugin-main/main.go:${PREFIX}/bin/session-manager-plugin
GO_TARGET+=	src/ssmcli-main/main.go:${PREFIX}/bin/ssmcli

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PORTDOCS=	NOTICE README.md RELEASENOTES.md THIRD-PARTY VERSION
LOGDIR=		${DESTDIR}/var/log/${PORTNAME}

OPTIONS_DEFINE=	DOCS

PLIST_FILES=	${ETCDIR}/sessionmanagerplugin/seelog.xml.template \
		${ETCDIR}/SSMCLI/seelog.xml.template \
		bin/session-manager-plugin \
		bin/ssmcli \
		'@dir /var/log/amazon-ssm-plugin'

post-patch: .PHONY .SILENT
	${CP} ${FILESDIR}/go.mod ${WRKSRC}
	${CP} ${FILESDIR}/go.sum ${WRKSRC}
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

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${CP} ${FILESDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY .SILENT
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME}
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy
	${TAR} --create --directory ${FILESDIR} --exclude=go.mod --exclude=go.sum --xz --file ${TMPDIR}/${DISTNAME}.tar.xz .
	${FIND} ${FILESDIR} -type f -delete
	${CP} ${WRKSRC}/go.mod ${FILESDIR}
	${CP} ${WRKSRC}/go.sum ${FILESDIR}

pre-deps-all: .PHONY
	${TAR} --extract --directory ${FILESDIR} --file ${TMPDIR}/${DISTNAME}.tar.xz
	${RM} ${PATCH_COOKIE} ${TMPDIR}/${DISTNAME}.tar.xz
	${CP} ${WRKSRC}/vendor/modules.txt ${FILESDIR}

post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-twinj:uuid:v1.0.0:twinj_uuid
GH_TUPLE:=	${GH_TUPLE:Ntwinj\:uuid\:*}

.include <bsd.port.mk>
