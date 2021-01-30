package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

//func init() {
//	time.Sleep(time.Duration(2) * time.Second)
//	data.RunName = "hhhh"
//}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {"Home", "", welcomeScreen},
		"binding": {"Add Web Server",
			"Don't Input Chinese",
			bindingScreen},
		"animations": {"Animations",
			"See how to animate components.",
			makeAnimationScreen,
		},
		"webBuild": {"Web Auto Publish",
			"Please select the project root file address",
			webBuild,
		},
		"containers": {"Containers",
			"Containers group other widgets and canvas objects, organising according to their layout.\n" +
				"Standard containers are illustrated in this section, but developers can also provide custom " +
				"layouts using the fyne.NewContainerWithLayout() constructor.",
			containerScreen,
		},
		"apptabs": {"AppTabs",
			"A container to help divide up an application into functional areas.",
			makeAppTabsTab,
		},
		"border": {"Border",
			"A container that positions items around a central content.",
			makeBorderLayout,
		},
		"box": {"Box",
			"A container arranges items in horizontal or vertical list.",
			makeBoxLayout,
		},
		"center": {"Center",
			"A container to that centers child elements.",
			makeCenterLayout,
		},
		"grid": {"Grid",
			"A container that arranges all items in a grid.",
			makeGridLayout,
		},
		"split": {"Split",
			"A split container divides the container in two pieces that the user can resize.",
			makeSplitTab,
		},
		"scroll": {"Scroll",
			"A container that provides scrolling for it's content.",
			makeScrollTab,
		},
		"widgets": {"Widgets",
			"In this section you can see the features available in the toolkit widget set.\n" +
				"Expand the tree on the left to browse the individual tutorial elements.",
			widgetScreen,
		},
		"accordion": {"Accordion",
			"Expand or collapse content panels.",
			makeAccordionTab,
		},
		"button": {"Button",
			"Simple widget for user tap handling.",
			makeButtonTab,
		},
		"card": {"Card",
			"Group content and widgets.",
			makeCardTab,
		},
		"entry": {"Entry",
			"Different ways to use the entry widget.",
			makeEntryTab,
		},
		"form": {"Form",
			"Gathering input widgets for data submission.",
			makeFormTab,
		},
		"input": {"Input",
			"A collection of widgets for user input.",
			makeInputTab,
		},
		"text": {"Text",
			"Text handling widgets.",
			makeTextTab,
		},
		"toolbar": {"Toolbar",
			"A row of shortcut icons for common tasks.",
			makeToolbarTab,
		},
		"progress": {"Progress",
			"Show duration or the need to wait for a task.",
			makeProgressTab,
		},
		"collections": {"Collections",
			"Collection widgets provide an efficient way to present lots of content.\n" +
				"The List, Table, and Tree provide a cache and re-use mechanism that make it possible to scroll through thousands of elements.\n" +
				"Use this for large data sets or for collections that can expand as users scroll.",
			collectionScreen,
		},
		"list": {"List",
			"A vertical arrangement of cached elements with the same styling.",
			makeListTab,
		},
		"table": {"Table",
			"A two dimensional cached collection of cells.",
			makeTableTab,
		},
		"tree": {"Tree",
			"A tree based arrangement of cached elements with the same styling.",
			makeTreeTab,
		},
		"dialogs": {"Dialogs",
			"Work with dialogs.",
			dialogScreen,
		},
		"windows": {"Change Server Provider",
			"Select a Provider and Click Run",
			windowScreen,
		},
		"advanced": {"Advanced",
			"Debug and advanced information.",
			advancedScreen,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "windows", "binding", "webBuild"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "grid", "split", "scroll"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "text", "toolbar", "progress"},
	}
)
