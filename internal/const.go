// Copyright 2022 Google LLC
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 2 as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// Contents of this file are derived from gpib_user.h in linux-gpib-4.1.0,
// which is copyright 2002 by Frank Mori Hess, fmhess@users.sourceforge.net.

package internal

/*
#include <gpib/ib.h>
*/
import "C"

const (
	GPIB_MAX_NUM_BOARDS      = C.GPIB_MAX_NUM_BOARDS
	GPIB_MAX_NUM_DESCRIPTORS = C.GPIB_MAX_NUM_DESCRIPTORS

	/* IBSTA status bits (returned by all functions) */
	DCAS  = C.DCAS  /* device clear state */
	DTAS  = C.DTAS  /* device trigger state */
	LACS  = C.LACS  /* GPIB interface is addressed as Listener */
	TACS  = C.TACS  /* GPIB interface is addressed as Talker */
	ATN   = C.ATN   /* Attention is asserted */
	CIC   = C.CIC   /* GPIB interface is Controller-in-Charge */
	REM   = C.REM   /* remote state */
	LOK   = C.LOK   /* lockout state */
	CMPL  = C.CMPL  /* I/O is complete */
	EVENT = C.EVENT /* DCAS, DTAS, or IFC has occurred */
	SPOLL = C.SPOLL /* board serial polled by busmaster */
	RQS   = C.RQS   /* Device requesting service */
	SRQI  = C.SRQI  /* SRQ is asserted */
	END   = C.END   /* EOI or EOS encountered */
	TIMO  = C.TIMO  /* Time limit on I/O or wait function exceeded */
	ERR   = C.ERR   /* Function call terminated on error */

	/* IBERR error codes */
	EDVR = C.EDVR /* system error */
	ECIC = C.ECIC /* not CIC */
	ENOL = C.ENOL /* no listeners */
	EADR = C.EADR /* CIC and not addressed before I/O */
	EARG = C.EARG /* bad argument to function call */
	ESAC = C.ESAC /* not SAC */
	EABO = C.EABO /* I/O operation was aborted */
	ENEB = C.ENEB /* non-existent board (GPIB interface offline) */
	EDMA = C.EDMA /* DMA hardware error detected */
	EOIP = C.EOIP /* new I/O attempted with old I/O in progress */
	ECAP = C.ECAP /* no capability for intended opeation */
	EFSO = C.EFSO /* file system operation error */
	EBUS = C.EBUS /* bus error */
	ESTB = C.ESTB /* lost serial poll bytes */
	ESRQ = C.ESRQ /* SRQ stuck on */
	ETAB = C.ETAB /* Table Overflow */

	/* Timeout values and meanings */
	TNONE  = C.TNONE  /* Infinite timeout (disabled) */
	T10us  = C.T10us  /* Timeout of 10 usec (ideal) */
	T30us  = C.T30us  /* Timeout of 30 usec (ideal) */
	T100us = C.T100us /* Timeout of 100 usec (ideal) */
	T300us = C.T300us /* Timeout of 300 usec (ideal) */
	T1ms   = C.T1ms   /* Timeout of 1 msec (ideal) */
	T3ms   = C.T3ms   /* Timeout of 3 msec (ideal) */
	T10ms  = C.T10ms  /* Timeout of 10 msec (ideal) */
	T30ms  = C.T30ms  /* Timeout of 30 msec (ideal) */
	T100ms = C.T100ms /* Timeout of 100 msec (ideal) */
	T300ms = C.T300ms /* Timeout of 300 msec (ideal) */
	T1s    = C.T1s    /* Timeout of 1 sec (ideal) */
	T3s    = C.T3s    /* Timeout of 3 sec (ideal) */
	T10s   = C.T10s   /* Timeout of 10 sec (ideal) */
	T30s   = C.T30s   /* Timeout of 30 sec (ideal) */
	T100s  = C.T100s  /* Timeout of 100 sec (ideal) */
	T300s  = C.T300s  /* Timeout of 300 sec (ideal) */
	T1000s = C.T1000s /* Timeout of 1000 sec (maximum) */

	/* End-of-string (EOS) modes for use with ibeos */
	EOS_MASK = C.EOS_MASK
	REOS     = C.REOS /* Terminate reads on EOS */
	XEOS     = C.XEOS /* assert EOI when EOS char is sent */
	BIN      = C.BIN  /* Do 8-bit compare on EOS */

	/* GPIB Bus Control Lines bit vector */
	ValidDAV  = C.ValidDAV
	ValidNDAC = C.ValidNDAC
	ValidNRFD = C.ValidNRFD
	ValidIFC  = C.ValidIFC
	ValidREN  = C.ValidREN
	ValidSRQ  = C.ValidSRQ
	ValidATN  = C.ValidATN
	ValidEOI  = C.ValidEOI
	ValidALL  = C.ValidALL
	BusDAV    = C.BusDAV  /* DAV  line status bit */
	BusNDAC   = C.BusNDAC /* NDAC line status bit */
	BusNRFD   = C.BusNRFD /* NRFD line status bit */
	BusIFC    = C.BusIFC  /* IFC  line status bit */
	BusREN    = C.BusREN  /* REN  line status bit */
	BusSRQ    = C.BusSRQ  /* SRQ  line status bit */
	BusATN    = C.BusATN  /* ATN  line status bit */
	BusEOI    = C.BusEOI  /* EOI  line status bit */

	/* Possible GPIB command messages */
	GTL = C.GTL /* go to local */
	SDC = C.SDC /* selected device clear */
	PPC = C.PPC /* parallel poll configure */
	GET = C.GET /* group execute trigger */
	TCT = C.TCT /* take control */
	LLO = C.LLO /* local lockout */
	DCL = C.DCL /* device clear */
	PPU = C.PPU /* parallel poll unconfigure */
	SPE = C.SPE /* serial poll enable */
	SPD = C.SPD /* serial poll disable */
	LAD = C.LAD /* value to be 'ored' in to obtain listen address */
	UNL = C.UNL /* unlisten */
	TAD = C.TAD /* value to be 'ored' in to obtain talk address */
	UNT = C.UNT /* untalk */
	SAD = C.SAD /* my secondary address (base) */
	PPE = C.PPE /* parallel poll enable (base) */
	PPD = C.PPD /* parallel poll disable */

	/* ppe_bits */
	PPC_DISABLE  = C.PPC_DISABLE
	PPC_SENSE    = C.PPC_SENSE /* parallel poll sense bit */
	PPC_DIO_MASK = C.PPC_DIO_MASK

	/* ibask_option */
	IbaPAD            = C.IbaPAD
	IbaSAD            = C.IbaSAD
	IbaTMO            = C.IbaTMO
	IbaEOT            = C.IbaEOT
	IbaPPC            = C.IbaPPC      /* board only */
	IbaREADDR         = C.IbaREADDR   /* device only */
	IbaAUTOPOLL       = C.IbaAUTOPOLL /* board only */
	IbaCICPROT        = C.IbaCICPROT  /* board only */
	IbaIRQ            = C.IbaIRQ      /* board only */
	IbaSC             = C.IbaSC       /* board only */
	IbaSRE            = C.IbaSRE      /* board only */
	IbaEOSrd          = C.IbaEOSrd
	IbaEOSwrt         = C.IbaEOSwrt
	IbaEOScmp         = C.IbaEOScmp
	IbaEOSchar        = C.IbaEOSchar
	IbaPP2            = C.IbaPP2    /* board only */
	IbaTIMING         = C.IbaTIMING /* board only */
	IbaDMA            = C.IbaDMA    /* board only */
	IbaReadAdjust     = C.IbaReadAdjust
	IbaWriteAdjust    = C.IbaWriteAdjust
	IbaEventQueue     = C.IbaEventQueue /* board only */
	IbaSPollBit       = C.IbaSPollBit   /* board only */
	IbaSpollBit       = C.IbaSpollBit   /* board only */
	IbaSendLLO        = C.IbaSendLLO    /* board only */
	IbaSPollTime      = C.IbaSPollTime  /* device only */
	IbaPPollTime      = C.IbaPPollTime  /* board only */
	IbaEndBitIsNormal = C.IbaEndBitIsNormal
	IbaUnAddr         = C.IbaUnAddr        /* device only */
	IbaHSCableLength  = C.IbaHSCableLength /* board only */
	IbaIst            = C.IbaIst           /* board only */
	IbaRsv            = C.IbaRsv           /* board only */
	IbaBNA            = C.IbaBNA           /* device only */
	/* linux-gpib extensions */
	Iba7BitEOS = C.Iba7BitEOS /* board only. Returns 1 if board supports 7 bit eos compares*/

	/* ibconfig_option */
	IbcPAD            = C.IbcPAD
	IbcSAD            = C.IbcSAD
	IbcTMO            = C.IbcTMO
	IbcEOT            = C.IbcEOT
	IbcPPC            = C.IbcPPC      /* board only */
	IbcREADDR         = C.IbcREADDR   /* device only */
	IbcAUTOPOLL       = C.IbcAUTOPOLL /* board only */
	IbcCICPROT        = C.IbcCICPROT  /* board only */
	IbcIRQ            = C.IbcIRQ      /* board only */
	IbcSC             = C.IbcSC       /* board only */
	IbcSRE            = C.IbcSRE      /* board only */
	IbcEOSrd          = C.IbcEOSrd
	IbcEOSwrt         = C.IbcEOSwrt
	IbcEOScmp         = C.IbcEOScmp
	IbcEOSchar        = C.IbcEOSchar
	IbcPP2            = C.IbcPP2    /* board only */
	IbcTIMING         = C.IbcTIMING /* board only */
	IbcDMA            = C.IbcDMA    /* board only */
	IbcReadAdjust     = C.IbcReadAdjust
	IbcWriteAdjust    = C.IbcWriteAdjust
	IbcEventQueue     = C.IbcEventQueue /* board only */
	IbcSPollBit       = C.IbcSPollBit   /* board only */
	IbcSpollBit       = C.IbcSpollBit   /* board only */
	IbcSendLLO        = C.IbcSendLLO    /* board only */
	IbcSPollTime      = C.IbcSPollTime  /* device only */
	IbcPPollTime      = C.IbcPPollTime  /* board only */
	IbcEndBitIsNormal = C.IbcEndBitIsNormal
	IbcUnAddr         = C.IbcUnAddr        /* device only */
	IbcHSCableLength  = C.IbcHSCableLength /* board only */
	IbcIst            = C.IbcIst           /* board only */
	IbcRsv            = C.IbcRsv           /* board only */
	IbcBNA            = C.IbcBNA           /* device only */

	/* t1_delays */
	T1_DELAY_2000ns = C.T1_DELAY_2000ns
	T1_DELAY_500ns  = C.T1_DELAY_500ns
	T1_DELAY_350ns  = C.T1_DELAY_350ns

	/* gpib_events */
	EventNone   = C.EventNone
	EventDevTrg = C.EventDevTrg
	EventDevClr = C.EventDevClr
	EventIFC    = C.EventIFC

	/* gpib_stb */
	IbStbRQS = C.IbStbRQS /* IEEE 488.1 & 2 */
	IbStbESB = C.IbStbESB /* IEEE 488.2 only */
	IbStbMAV = C.IbStbMAV /* IEEE 488.2 only */

	/* sad_special_address */
	NO_SAD  = C.NO_SAD
	ALL_SAD = C.ALL_SAD

	/* send_eotmode */
	NULLend = C.NULLend
	DABend  = C.DABend
	NLend   = C.NLend

	/* static constants from ib.h */
	NOADDR  = 0xffff
	STOPend = 0x100
)
