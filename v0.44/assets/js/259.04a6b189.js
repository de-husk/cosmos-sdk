(window.webpackJsonp=window.webpackJsonp||[]).push([[259],{696:function(e,a,t){"use strict";t.r(a);var r=t(1),s=Object(r.a)({},(function(){var e=this,a=e.$createElement,t=e._self._c||a;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"subspace"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#subspace"}},[e._v("#")]),e._v(" Subspace")]),e._v(" "),t("p",[t("code",[e._v("Subspace")]),e._v(" is a prefixed subspace of the parameter store. Each module which uses the\nparameter store will take a "),t("code",[e._v("Subspace")]),e._v(" to isolate permission to access.")]),e._v(" "),t("h2",{attrs:{id:"key"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#key"}},[e._v("#")]),e._v(" Key")]),e._v(" "),t("p",[e._v("Parameter keys are human readable alphanumeric strings. A parameter for the key\n"),t("code",[e._v('"ExampleParameter"')]),e._v(" is stored under "),t("code",[e._v('[]byte("SubspaceName" + "/" + "ExampleParameter")')]),e._v(",\nwhere "),t("code",[e._v('"SubspaceName"')]),e._v(" is the name of the subspace.")]),e._v(" "),t("p",[e._v("Subkeys are secondary parameter keys those are used along with a primary parameter key.\nSubkeys can be used for grouping or dynamic parameter key generation during runtime.")]),e._v(" "),t("h2",{attrs:{id:"keytable"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#keytable"}},[e._v("#")]),e._v(" KeyTable")]),e._v(" "),t("p",[e._v("All of the parameter keys that will be used should be registered at the compile\ntime. "),t("code",[e._v("KeyTable")]),e._v(" is essentially a "),t("code",[e._v("map[string]attribute")]),e._v(", where the "),t("code",[e._v("string")]),e._v(" is a parameter key.")]),e._v(" "),t("p",[e._v("Currently, "),t("code",[e._v("attribute")]),e._v(" consists of a "),t("code",[e._v("reflect.Type")]),e._v(", which indicates the parameter\ntype to check that provided key and value are compatible and registered, as well as a function "),t("code",[e._v("ValueValidatorFn")]),e._v(" to validate values.")]),e._v(" "),t("p",[e._v("Only primary keys have to be registered on the "),t("code",[e._v("KeyTable")]),e._v(". Subkeys inherit the\nattribute of the primary key.")]),e._v(" "),t("h2",{attrs:{id:"paramset"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#paramset"}},[e._v("#")]),e._v(" ParamSet")]),e._v(" "),t("p",[e._v("Modules often define parameters as a proto message. The generated struct can implement\n"),t("code",[e._v("ParamSet")]),e._v(" interface to be used with the following methods:")]),e._v(" "),t("ul",[t("li",[t("code",[e._v("KeyTable.RegisterParamSet()")]),e._v(": registers all parameters in the struct")]),e._v(" "),t("li",[t("code",[e._v("Subspace.{Get, Set}ParamSet()")]),e._v(": Get to & Set from the struct")])]),e._v(" "),t("p",[e._v("The implementor should be a pointer in order to use "),t("code",[e._v("GetParamSet()")]),e._v(".")])])}),[],!1,null,null,null);a.default=s.exports}}]);