<?php

namespace App\Http\Controllers;

use App\Http\Resources\EducationResource;
use Exception;
use App\Http\Resources\InformationResource;
use App\Models\Education;
use Illuminate\Http\Request;
use App\Models\Information;

class GeneralController extends Controller
{

    public function getInformation()
    {
        try {
            $information = new Information;
            return new InformationResource($information);
        } catch (Exception $e) {
            return response()->json(['error' => $e->getMessage()], 400);
        }
    }

    public function postEducation(Request $request)
    {
        $education = Education::find($request->id);

        if (empty($education)) {
            $education = new Education();
        }

        if ($request->has('degree')) {
            $degree = $request->degree;
            $education->degree = $degree;
        }

        if ($request->has('major')) {
            $major = $request->major;
            $education->major = $major;
        }

        if ($request->has('university_name')) {
            $universityName = $request->university_name;
            $education->university_name = $universityName;
        }

        if ($request->has('university_abrv')) {
            $universityAbrv = $request->university_abrv;
            $education->university_abrv = $universityAbrv;
        }

        if ($request->has('date')) {
            $date = $request->date;
            $education->date = $date;
        }

        $education->save();


        return new EducationResource($education);
    }
}
