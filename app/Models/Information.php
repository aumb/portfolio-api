<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Information extends Model
{
    public function jobs()
    {
        $jobs = Job::all();
        return $jobs;
    }

    public function education()
    {
        $education = Education::all();
        return $education->first();
    }

    public function recentWork()
    {
        $recentWork = RecentWork::all();
        return $recentWork;
    }

    public function personalInformation()
    {
        $personalInformation = PersonalInformation::all();
        return $personalInformation->first();
    }
}
