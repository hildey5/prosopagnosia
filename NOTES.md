# Links to use as examples 

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html

c:\Users\lucas\go\bin\build-lambda-zip.exe -o detectFaceHandler.zip detectFaceHandler

package main

import (
    "context"
    "fmt"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func lambda_handler(ctx context.Context, s3Event events.S3Event) {
    for _, record := range s3Event.Records {
        s3 := record.S3
        fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
    }
}

func main() {
    lambda.Start(lambda_handler)
}
GOOS=linux go build detectFaceHandler.go
go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-lambda-go/events
zip detectFaceHandler.zip detectFaceHandler



aws rekognition index-faces --image {\"S3Object\":{\"Bucket\":\"faces.collection\",\"Name\":\"tammy.jpg\"}} --collection-id "people_lucas" --max-faces 1 --quality-filter "AUTO" --detection-attributes "ALL" --external-image-id "tammy.jpg"


{
    "FaceRecords": [
        {
            "Face": {
                "FaceId": "d4d40d8b-6b7d-48d7-831d-bc0a483309d5",
                "BoundingBox": {
                    "Width": 0.29369065165519714,
                    "Height": 0.3033809959888458,
                    "Left": 0.25295037031173706,
                    "Top": 0.3832453787326813
                },
                "ImageId": "6ffb2aa9-d711-3581-8a6a-19d12540acdd",
                "ExternalImageId": "tammy.jpg",
                "Confidence": 100.0
            },
            "FaceDetail": {
                "BoundingBox": {
                    "Width": 0.29369065165519714,
                    "Height": 0.3033809959888458,
                    "Left": 0.25295037031173706,
                    "Top": 0.3832453787326813
                },
                "AgeRange": {
                    "Low": 27,
                    "High": 44
                },
                "Smile": {
                    "Value": true,
                    "Confidence": 99.98711395263672
                },
                "Eyeglasses": {
                    "Value": false,
                    "Confidence": 99.98077392578125
                },
                "Sunglasses": {
                    "Value": false,
                    "Confidence": 99.99567413330078
                },
                "Gender": {
                    "Value": "Female",
                    "Confidence": 99.9916000366211
                },
                "Beard": {
                    "Value": false,
                    "Confidence": 99.9990005493164
                },
                "Mustache": {
                    "Value": false,
                    "Confidence": 99.99872589111328
                },
                "EyesOpen": {
                    "Value": true,
                    "Confidence": 99.99637603759766
                },
                "MouthOpen": {
                    "Value": true,
                    "Confidence": 99.87557220458984
                },
                "Emotions": [
                    {
                        "Type": "CONFUSED",
                        "Confidence": 0.5347237586975098
                    },
                    {
                        "Type": "HAPPY",
                        "Confidence": 96.6146011352539
                    },
                    {
                        "Type": "ANGRY",
                        "Confidence": 0.47298464179039
                    },
                    {
                        "Type": "CALM",
                        "Confidence": 0.1840682029724121
                    },
                    {
                        "Type": "SAD",
                        "Confidence": 0.1390680968761444
                    },
                    {
                        "Type": "SURPRISED",
                        "Confidence": 1.4539893865585327
                    },
                    {
                        "Type": "DISGUSTED",
                        "Confidence": 0.6005598902702332
                    }
                ],
                "Landmarks": [
                    {
                        "Type": "eyeLeft",
                        "X": 0.33517342805862427,
                        "Y": 0.5025562047958374
                    },
                    {
                        "Type": "eyeRight",
                        "X": 0.4693393409252167,
                        "Y": 0.5064172744750977
                    },
                    {
                        "Type": "mouthLeft",
                        "X": 0.3426228165626526,
                        "Y": 0.6039745211601257
                    },
                    {
                        "Type": "mouthRight",
                        "X": 0.4541293978691101,
                        "Y": 0.6073617935180664
                    },
                    {
                        "Type": "nose",
                        "X": 0.39806586503982544,
                        "Y": 0.5590881705284119
                    },
                    {
                        "Type": "leftEyeBrowLeft",
                        "X": 0.28547415137290955,
                        "Y": 0.47717198729515076
                    },
                    {
                        "Type": "leftEyeBrowRight",
                        "X": 0.3632919192314148,
                        "Y": 0.4734383821487427
                    },
                    {
                        "Type": "leftEyeBrowUp",
                        "X": 0.3249589800834656,
                        "Y": 0.4666908383369446
                    },
                    {
                        "Type": "rightEyeBrowLeft",
                        "X": 0.44192805886268616,
                        "Y": 0.4756656587123871
                    },
                    {
                        "Type": "rightEyeBrowRight",
                        "X": 0.5236599445343018,
                        "Y": 0.48421722650527954
                    },
                    {
                        "Type": "rightEyeBrowUp",
                        "X": 0.4825075566768646,
                        "Y": 0.4714040458202362
                    },
                    {
                        "Type": "leftEyeLeft",
                        "X": 0.3127691149711609,
                        "Y": 0.500826358795166
                    },
                    {
                        "Type": "leftEyeRight",
                        "X": 0.3615882396697998,
                        "Y": 0.5036354660987854
                    },
                    {
                        "Type": "leftEyeUp",
                        "X": 0.33556583523750305,
                        "Y": 0.49723437428474426
                    },
                    {
                        "Type": "leftEyeDown",
                        "X": 0.336160808801651,
                        "Y": 0.5064941644668579
                    },
                    {
                        "Type": "rightEyeLeft",
                        "X": 0.44089868664741516,
                        "Y": 0.5058059096336365
                    },
                    {
                        "Type": "rightEyeRight",
                        "X": 0.49031591415405273,
                        "Y": 0.5056743025779724
                    },
                    {
                        "Type": "rightEyeUp",
                        "X": 0.4672227203845978,
                        "Y": 0.5008348226547241
                    },
                    {
                        "Type": "rightEyeDown",
                        "X": 0.4659702777862549,
                        "Y": 0.5100512504577637
                    },
                    {
                        "Type": "noseLeft",
                        "X": 0.3733530342578888,
                        "Y": 0.569662868976593
                    },
                    {
                        "Type": "noseRight",
                        "X": 0.4227769672870636,
                        "Y": 0.5705262422561646
                    },
                    {
                        "Type": "mouthUp",
                        "X": 0.3968857228755951,
                        "Y": 0.5944077372550964
                    },
                    {
                        "Type": "mouthDown",
                        "X": 0.39584818482398987,
                        "Y": 0.624567985534668
                    },
                    {
                        "Type": "leftPupil",
                        "X": 0.33517342805862427,
                        "Y": 0.5025562047958374
                    },
                    {
                        "Type": "rightPupil",
                        "X": 0.4693393409252167,
                        "Y": 0.5064172744750977
                    },
                    {
                        "Type": "upperJawlineLeft",
                        "X": 0.2551567554473877,
                        "Y": 0.4997469484806061
                    },
                    {
                        "Type": "midJawlineLeft",
                        "X": 0.2796677350997925,
                        "Y": 0.6113449335098267
                    },
                    {
                        "Type": "chinBottom",
                        "X": 0.39451664686203003,
                        "Y": 0.6773495674133301
                    },
                    {
                        "Type": "midJawlineRight",
                        "X": 0.5191413164138794,
                        "Y": 0.6176446080207825
                    },
                    {
                        "Type": "upperJawlineRight",
                        "X": 0.5534681677818298,
                        "Y": 0.507597029209137
                    }
                ],
                "Pose": {
                    "Roll": -3.224494457244873,
                    "Yaw": -15.90917682647705,
                    "Pitch": 8.298050880432129
                },
                "Quality": {
                    "Brightness": 63.65409469604492,
                    "Sharpness": 83.14741516113281
                },
                "Confidence": 100.0
            }
        }
    ],
    "FaceModelVersion": "4.0",
    "UnindexedFaces": []
}




C:\Users\lucas>aws rekognition index-faces --image {\"S3Object\":{\"Bucket\":\"faces.collection\",\"Name\":\"addison.jpg\"}} --collection-id "people_lucas" --max-faces 1 --quality-filter "AUTO" --detection-attributes "ALL" --external-image-id "addison.jpg"
{
    "FaceRecords": [
        {
            "Face": {
                "FaceId": "46339d0c-e224-4f5d-8d97-39e54256dbb2",
                "BoundingBox": {
                    "Width": 0.3233455419540405,
                    "Height": 0.331185907125473,
                    "Left": 0.3267185091972351,
                    "Top": 0.32170370221138
                },
                "ImageId": "bafac079-97f2-3d44-ae70-8b5f4b7f7c1d",
                "ExternalImageId": "addison.jpg",
                "Confidence": 100.0
            },
            "FaceDetail": {
                "BoundingBox": {
                    "Width": 0.3233455419540405,
                    "Height": 0.331185907125473,
                    "Left": 0.3267185091972351,
                    "Top": 0.32170370221138
                },
                "AgeRange": {
                    "Low": 10,
                    "High": 15
                },
                "Smile": {
                    "Value": true,
                    "Confidence": 98.54743194580078
                },
                "Eyeglasses": {
                    "Value": false,
                    "Confidence": 99.95611572265625
                },
                "Sunglasses": {
                    "Value": false,
                    "Confidence": 99.98686981201172
                },
                "Gender": {
                    "Value": "Female",
                    "Confidence": 99.8141098022461
                },
                "Beard": {
                    "Value": false,
                    "Confidence": 99.98085021972656
                },
                "Mustache": {
                    "Value": false,
                    "Confidence": 99.99267578125
                },
                "EyesOpen": {
                    "Value": true,
                    "Confidence": 99.99878692626953
                },
                "MouthOpen": {
                    "Value": false,
                    "Confidence": 97.68605041503906
                },
                "Emotions": [
                    {
                        "Type": "DISGUSTED",
                        "Confidence": 0.0459308996796608
                    },
                    {
                        "Type": "HAPPY",
                        "Confidence": 98.07469177246094
                    },
                    {
                        "Type": "ANGRY",
                        "Confidence": 0.03324158117175102
                    },
                    {
                        "Type": "CONFUSED",
                        "Confidence": 0.10261669754981995
                    },
                    {
                        "Type": "CALM",
                        "Confidence": 1.599405288696289
                    },
                    {
                        "Type": "SURPRISED",
                        "Confidence": 0.10466409474611282
                    },
                    {
                        "Type": "SAD",
                        "Confidence": 0.039441369473934174
                    }
                ],
                "Landmarks": [
                    {
                        "Type": "eyeLeft",
                        "X": 0.4146278202533722,
                        "Y": 0.4654943346977234
                    },
                    {
                        "Type": "eyeRight",
                        "X": 0.5611647963523865,
                        "Y": 0.4656446576118469
                    },
                    {
                        "Type": "mouthLeft",
                        "X": 0.42903971672058105,
                        "Y": 0.5756911635398865
                    },
                    {
                        "Type": "mouthRight",
                        "X": 0.5506148934364319,
                        "Y": 0.5759022235870361
                    },
                    {
                        "Type": "nose",
                        "X": 0.48838940262794495,
                        "Y": 0.5245154500007629
                    },
                    {
                        "Type": "leftEyeBrowLeft",
                        "X": 0.3587990403175354,
                        "Y": 0.43893396854400635
                    },
                    {
                        "Type": "leftEyeBrowRight",
                        "X": 0.4442847669124603,
                        "Y": 0.4326961934566498
                    },
                    {
                        "Type": "leftEyeBrowUp",
                        "X": 0.4019632339477539,
                        "Y": 0.42625874280929565
                    },
                    {
                        "Type": "rightEyeBrowLeft",
                        "X": 0.5301116704940796,
                        "Y": 0.43273112177848816
                    },
                    {
                        "Type": "rightEyeBrowRight",
                        "X": 0.6187273263931274,
                        "Y": 0.439902126789093
                    },
                    {
                        "Type": "rightEyeBrowUp",
                        "X": 0.5738422274589539,
                        "Y": 0.42698729038238525
                    },
                    {
                        "Type": "leftEyeLeft",
                        "X": 0.3903031647205353,
                        "Y": 0.46385207772254944
                    },
                    {
                        "Type": "leftEyeRight",
                        "X": 0.4439449608325958,
                        "Y": 0.4653535783290863
                    },
                    {
                        "Type": "leftEyeUp",
                        "X": 0.4152067005634308,
                        "Y": 0.4592190980911255
                    },
                    {
                        "Type": "leftEyeDown",
                        "X": 0.41639459133148193,
                        "Y": 0.4692700207233429
                    },
                    {
                        "Type": "rightEyeLeft",
                        "X": 0.5305822491645813,
                        "Y": 0.465302050113678
                    },
                    {
                        "Type": "rightEyeRight",
                        "X": 0.5842585563659668,
                        "Y": 0.4636955261230469
                    },
                    {
                        "Type": "rightEyeUp",
                        "X": 0.5590155124664307,
                        "Y": 0.45910415053367615
                    },
                    {
                        "Type": "rightEyeDown",
                        "X": 0.5581867098808289,
                        "Y": 0.46915796399116516
                    },
                    {
                        "Type": "noseLeft",
                        "X": 0.4611159563064575,
                        "Y": 0.5370980501174927
                    },
                    {
                        "Type": "noseRight",
                        "X": 0.5152450799942017,
                        "Y": 0.5362125635147095
                    },
                    {
                        "Type": "mouthUp",
                        "X": 0.48860061168670654,
                        "Y": 0.5630020499229431
                    },
                    {
                        "Type": "mouthDown",
                        "X": 0.4891490638256073,
                        "Y": 0.5957944989204407
                    },
                    {
                        "Type": "leftPupil",
                        "X": 0.4146278202533722,
                        "Y": 0.4654943346977234
                    },
                    {
                        "Type": "rightPupil",
                        "X": 0.5611647963523865,
                        "Y": 0.4656446576118469
                    },
                    {
                        "Type": "upperJawlineLeft",
                        "X": 0.3248665928840637,
                        "Y": 0.4646230638027191
                    },
                    {
                        "Type": "midJawlineLeft",
                        "X": 0.3585464358329773,
                        "Y": 0.5855122208595276
                    },
                    {
                        "Type": "chinBottom",
                        "X": 0.49002504348754883,
                        "Y": 0.653557300567627
                    },
                    {
                        "Type": "midJawlineRight",
                        "X": 0.6205476522445679,
                        "Y": 0.5852450132369995
                    },
                    {
                        "Type": "upperJawlineRight",
                        "X": 0.6512303352355957,
                        "Y": 0.4642896354198456
                    }
                ],
                "Pose": {
                    "Roll": -5.4465718269348145,
                    "Yaw": -14.79128646850586,
                    "Pitch": 8.318347930908203
                },
                "Quality": {
                    "Brightness": 52.78631591796875,
                    "Sharpness": 83.14741516113281
                },
                "Confidence": 100.0
            }
        }
    ],
    "FaceModelVersion": "4.0",
    "UnindexedFaces": []
}




C:\Users\lucas>aws rekognition search-faces-by-image --image {\"S3Object\":{\"Bucket\":\"faces.collection\",\"Name\":\"addison.jpg\"}} --collection-id "people_lucas"
{
    "SearchedFaceBoundingBox": {
        "Width": 0.3233455419540405,
        "Height": 0.331185907125473,
        "Left": 0.3267185091972351,
        "Top": 0.32170370221138
    },
    "SearchedFaceConfidence": 100.0,
    "FaceMatches": [
        {
            "Similarity": 99.99999237060547,
            "Face": {
                "FaceId": "46339d0c-e224-4f5d-8d97-39e54256dbb2",
                "BoundingBox": {
                    "Width": 0.32334598898887634,
                    "Height": 0.3311859965324402,
                    "Left": 0.3267189860343933,
                    "Top": 0.3217040002346039
                },
                "ImageId": "bafac079-97f2-3d44-ae70-8b5f4b7f7c1d",
                "ExternalImageId": "addison.jpg",
                "Confidence": 100.0
            }
        }
    ],
    "FaceModelVersion": "4.0"
}