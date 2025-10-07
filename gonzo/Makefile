PORTNAME=	gonzo
PORTVERSION=	0.2.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	TUI log analysis tool
WWW=		https://gonzo.controltheory.com

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	control-theory

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-charmbracelet:x:ansi/v0.9.3:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x
#+charmbracelet:x:ansi/v0.9.3:charmbracelet_x_ansi
# charmbracelet:x:cellbuf/v0.0.13:charmbracelet_x_cellbuf
# charmbracelet:x:term/v0.2.1:charmbracelet_x_term
GH_TUPLE:=	${GH_TUPLE:C,(:charmbracelet_x_(ansi|cellbuf|term)).*,\1,}

#-google:go-genproto:200df99c418a:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:200df99c418a:google_go_genproto_api
#-google:go-genproto:5f3141c8851a:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:5f3141c8851a:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C,(:google_go_genproto).*/(api|rpc),\1_\2,}

#-open-telemetry:opentelemetry-proto-go:v1.7.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto
#+open-telemetry:opentelemetry-proto-go:v1.7.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
GH_TUPLE:=	${GH_TUPLE:S,/otlp$,,1}

.include <bsd.port.mk>
