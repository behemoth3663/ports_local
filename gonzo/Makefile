PORTNAME=	gonzo
PORTVERSION=	0.4.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	TUI log analysis tool
WWW=		https://gonzo.controltheory.com

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	control-theory

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-charmbracelet:x:ansi/v0.9.3:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x
#+charmbracelet:x:ansi/v0.9.3:charmbracelet_x_ansi
# charmbracelet:x:cellbuf/v0.0.13:charmbracelet_x_cellbuf
# charmbracelet:x:term/v0.2.1:charmbracelet_x_term
GH_TUPLE:=	${GH_TUPLE:C,(:charmbracelet_x_(ansi|cellbuf|term)).*,\1,}

#-charmbracelet:x:2fdc97757edf:charmbracelet_x_slice
#+charmbracelet:x:2fdc97757edf:charmbracelet_x_slice/vendor/github.com/charmbracelet/x/exp/slice
GH_TUPLE:=	${GH_TUPLE:S,:charmbracelet_x_slice$,&/vendor/github.com/charmbracelet/x,1}

#-open-telemetry:opentelemetry-proto-go:v1.7.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto
#+open-telemetry:opentelemetry-proto-go:v1.7.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
GH_TUPLE:=	${GH_TUPLE:S,/otlp$,,1}

.include <bsd.port.mk>
