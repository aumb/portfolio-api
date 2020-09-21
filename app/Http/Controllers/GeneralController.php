<?php

namespace App\Http\Controllers;

use App\Http\Resources\EducationResource;
use Exception;
use App\Http\Resources\InformationResource;
use App\Models\Education;
use App\Models\PersonalInformation;
use Illuminate\Http\Request;
use App\Models\Information;
use Illuminate\Support\Carbon;

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
            $education->end_date = Carbon::createFromFormat('Y-m-d', $date);
        }

        $education->save();


        return new EducationResource($education);
    }

    public function postPersonalInformation(Request $request)
    {
        $personalInformation = PersonalInformation::find($request->id);

        if (empty($personalInformation)) {
            $personalInformation = new PersonalInformation();
        }

        if ($request->has('name')) {
            $name = $request->name;
            $personalInformation->name = $name;
        }

        if ($request->has('job_title')) {
            $jobTitle = $request->job_title;
            $personalInformation->job_title = $jobTitle;
        }

        if ($request->has('about')) {
            $about = $request->about;
            $personalInformation->about = $about;
        }

        if ($request->has('email')) {
            $email = $request->email;
            $personalInformation->email = $email;
        }

        if ($request->has('phone_number')) {
            $phoneNumber = $request->phone_number;
            $personalInformation->phone_number = $phoneNumber;
        }

        if ($request->has('linked_in_url')) {
            $linkedInUrl = $request->linked_in_url;
            $personalInformation->linked_in_url = $linkedInUrl;
        }

        if ($request->has('facebook_url')) {
            $facebookUrl = $request->facebook_url;
            $personalInformation->facebook_url = $facebookUrl;
        }

        if ($request->has('instagram_url')) {
            $instagramUrl = $request->instagram_url;
            $personalInformation->instagram_url = $instagramUrl;
        }

        if ($request->has('twitter_url')) {
            $twitterUrl = $request->twitter_url;
            $personalInformation->twitter_url = $twitterUrl;
        }

        if ($request->has('github_url')) {
            $githubUrl = $request->github_url;
            $personalInformation->github_url = $githubUrl;
        }

        if ($request->has('dob')) {
            $date = $request->dob;
            $personalInformation->dob = Carbon::createFromFormat('Y-m-d', $date);
        }

        $personalInformation->save();


        return new PersonalInformation($personalInformation);
    }
}
