<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;
use App\Http\Resources\JobResourceCollection;
use App\Http\Resources\PersonalInformationResource;
use App\Http\Resources\EducationResource;

class InformationResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function toArray($request)
    {
        return [
            'jobs' => new JobResourceCollection($this->jobs()),
            'education' => new EducationResource($this->education()),
            'personal_information' => new PersonalInformationResource($this->personalInformation()),
        ];
    }
}
