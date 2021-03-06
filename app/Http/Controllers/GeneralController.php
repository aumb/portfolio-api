<?php

namespace App\Http\Controllers;

use App\Http\Resources\EducationResource;
use Exception;
use App\Http\Resources\InformationResource;
use App\Http\Resources\JobResource;
use App\Http\Resources\PersonalInformationResource;
use App\Http\Resources\RecentWorkResource;
use App\Models\Job;
use App\Models\Education;
use App\Models\PersonalInformation;
use Illuminate\Http\Request;
use App\Models\Information;
use App\Models\RecentWork;
use Illuminate\Support\Carbon;
use Illuminate\Support\Facades\Storage;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Auth;
use Laravel\Passport\Client as OClient;
use App\Models\User;

class GeneralController extends Controller
{
    public function login(Request $request)
    {

        $rules = [
            'email' => 'required',
            'password' => 'required',
        ];

        $this->validate($request, $rules);

        $email = request('email');

        $password = request('password');


        if (Auth::attempt(['email' => $email, 'password' => $password])) {
            $token  = $this->generateAccessToken($email, $password);
            $user = User::where('email', $email)->first();
            return response()->json(['token' => $token, 'user' => $user], 200,);
        } else {
            return response()->json(['INVAlID'], 400,);
        }
    }

    public function register(Request $request)
    {

        $rules = [
            'name' => 'required',
            'email' => 'required|email|unique:users',
            'password' => 'required|min:6',
        ];

        $this->validate($request, $rules);

        $data = $request->all();

        $unencryptedPassword = $data['password'];

        $data['password'] = bcrypt($request->password);

        $user = User::create($data);

        $token  = $this->generateAccessToken($user->email, $unencryptedPassword);

        return response()->json(['token' => $token, 'user' => $user], 201,);
    }

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

        if ($request->has('job_subtitle')) {
            $jobSubitle = $request->job_subtitle;
            $personalInformation->job_subtitle = $jobSubitle;
        }

        if ($request->has('about_title')) {
            $aboutTitle = $request->about_title;
            $personalInformation->about_title = $aboutTitle;
        }

        if ($request->has('about')) {
            $about = $request->about;
            $personalInformation->about = $about;
        }

        if ($request->has('contact_title')) {
            $contactTitle = $request->contact_title;
            $personalInformation->contact_title = $contactTitle;
        }

        if ($request->has('contact_subtitle')) {
            $contactSubtitle = $request->contact_subtitle;
            $personalInformation->contact_subtitle = $contactSubtitle;
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

        if ($request->has('image')) {
            $file = $request->file('image');
            if (empty($file) && !empty($personalInformation->img)) {
                Storage::disk('s3')->delete($personalInformation->img);
            } else {
                $filename = 'pp-' . Carbon::now()->timestamp . '.' . $file->getClientOriginalExtension();

                if (!empty($personalInformation->img)) {
                    Storage::disk('s3')->delete($personalInformation->img);
                }

                $path = $file->storeAs('cv', $filename, 's3');
                $personalInformation->img = $path;
            }
        }

        $personalInformation->save();


        return new PersonalInformationResource($personalInformation);
    }

    public function getProfilePicture($personalInformationId)
    {

        $personalInformation = PersonalInformation::findOrFail($personalInformationId);


        if (empty($personalInformation->img)) {
            return null;
        } else {
            $response = Storage::disk('s3')->response($personalInformation->img);
            $response->headers->set('Content-Type', 'image/png');
            return $response;
        }
    }

    public function postJob(Request $request)
    {
        $job = Job::find($request->id);

        if (empty($job)) {
            $job = new Job();
        }

        if ($request->has('title')) {
            $title = $request->title;
            $job->title = $title;
        }

        if ($request->has('company_name')) {
            $companyName = $request->company_name;
            $job->company_name = $companyName;
        }

        if ($request->has('company_link')) {
            $companyLink = $request->company_link;
            $job->company_link = $companyLink;
        }

        if ($request->has('description')) {
            $description = $request->description;
            $job->description = $description;
        }

        if ($request->has('start_date')) {
            $startDate = $request->start_date;
            $job->start_date = Carbon::createFromFormat('Y-m-d', $startDate);
        }

        if ($request->has('end_date')) {
            $endDate = $request->end_date;
            $job->end_date = Carbon::createFromFormat('Y-m-d', $endDate);
        }

        if ($request->has('image')) {
            $file = $request->file('image');
            if (empty($file) && !empty($job->img)) {
                Storage::disk('s3')->delete($job->img);
            } else {
                $filename = 'job-' . Carbon::now()->timestamp . '.' . $file->getClientOriginalExtension();

                if (!empty($job->img)) {
                    Storage::disk('s3')->delete($job->img);
                }

                $path = $file->storeAs('cv', $filename, 's3');
                $job->img = $path;
            }
        }

        $job->save();


        return new JobResource($job);
    }

    public function postRecentWork(Request $request)
    {
        $recentWork = RecentWork::find($request->id);

        if (empty($recentWork)) {
            $recentWork = new RecentWork();
        }

        if ($request->has('url')) {
            $url = $request->url;
            $recentWork->url = $url;
        }

        if ($request->has('description')) {
            $description = $request->description;
            $recentWork->description = $description;
        }


        if ($request->has('image')) {
            $file = $request->file('image');
            if (empty($file) && !empty($recentWork->img)) {
                Storage::disk('s3')->delete($recentWork->img);
            } else {
                $filename = 'recent_work-' . Carbon::now()->timestamp . '.' . $file->getClientOriginalExtension();

                if (!empty($recentWork->img)) {
                    Storage::disk('s3')->delete($recentWork->img);
                }

                $path = $file->storeAs('cv', $filename, 's3');
                $recentWork->img = $path;
            }
        }

        $recentWork->save();


        return new RecentWorkResource($recentWork);
    }

    public function getRecentWorkPicture($recentWorkId)
    {

        $recentWork = RecentWork::findOrFail($recentWorkId);


        if (empty($recentWork->img)) {
            return null;
        } else {
            $response = Storage::disk('s3')->response($recentWork->img);
            $response->headers->set('Content-Type', 'image/png');
            return $response;
        }
    }

    public function getJobPicture($jobId)
    {

        $job = Job::findOrFail($jobId);


        if (empty($job->img)) {
            return null;
        } else {
            $response = Storage::disk('s3')->response($job->img);
            $response->headers->set('Content-Type', 'image/png');
            return $response;
        }
    }

    public function deleteRecentWork($recentWorkId)
    {
        $recentWork = RecentWork::findOrFail($recentWorkId);

        if (!empty($recentWork->img)) {
            Storage::disk('s3')->delete($recentWork->img);
        }

        $recentWork->delete();
        return response()->json(null, 200);
    }

    public function deletePersonalInformation($personalInformation)
    {
        $personalInformation = PersonalInformation::findOrFail($personalInformation);

        if (!empty($personalInformation->img)) {
            Storage::disk('s3')->delete($personalInformation->img);
        }

        $personalInformation->delete();
        return response()->json(null, 200);
    }

    public function deleteJob($jobId)
    {
        $job = Job::findOrFail($jobId);

        if (!empty($job->img)) {
            Storage::disk('s3')->delete($job->img);
        }

        $job->delete();
        return response()->json(null, 200);
    }

    private function generateAccessToken($email, $password)
    {
        $oClient = OClient::where('password_client', 1)->first();

        $http = new Client;

        $response = $http->post('https://floating-cove-36500.herokuapp.com/oauth/token', [
            'form_params' => [
                'grant_type' => 'password',
                'client_id' => $oClient->id,
                'client_secret' => $oClient->secret,
                'username' => $email,
                'password' => $password,
                'scope' => '*',
            ],
        ]);

        return json_decode((string) $response->getBody(), true);
    }
}
