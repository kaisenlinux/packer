
source "amazon-ebs" "ubuntu-1604" {
    int = 42
    nested_slice {
        dynamic "tag" {
            for_each = local.standard_tags
            content {
                key                 = tag.key
                value               = tag.value
            }
        }
    }
}

source "virtualbox-iso" "ubuntu-1204" {
    string   = "${source.name}-${source.type}"
    int      = 42
    int64    = 43
    bool     = true
    trilean  = true
    duration = "10s"

    map_string_string = {
        a = "b"
        c = "d"
    }
    slice_string = [
        "a",
        "b",
        "c",
    ]
    slice_slice_string = [
        ["a","b"],
        ["c","d"]
    ]

    data_source = data.amazon-ami.test.string

    nested {
        string   = "string"
        int      = 42
        int64    = 43
        bool     = true
        trilean  = true
        duration = "10s"
        map_string_string = {
            a = "b"
            c = "d"
        }
        slice_string = [
            "a",
            "b",
            "c",
        ]
        slice_slice_string = [
            ["a","b"],
            ["c","d"]
        ]
        data_source = data.amazon-ami.test.string
    }

    nested_slice {
        string   = "string"
        int      = 42
        int64    = 43
        bool     = true
        trilean  = true
        duration = "10s"
        map_string_string = {
            a = "b"
            c = "d"
        }
        slice_string = [
            "a",
            "b",
            "c",
        ]
        slice_slice_string = [
            ["a","b"],
            ["c","d"]
        ]
        data_source = data.amazon-ami.test.string
    }

    nested_slice {
        string   = "string"
        int      = 42
        int64    = 43
        bool     = true
        trilean  = true
        duration = "10s"
        map_string_string = {
            a = "b"
            c = "d"
        }
        slice_string = [
            "a",
            "b",
            "c",
        ]
        slice_slice_string = [
            ["a","b"],
            ["c","d"]
        ]
        data_source = data.amazon-ami.test.string
    }
}
