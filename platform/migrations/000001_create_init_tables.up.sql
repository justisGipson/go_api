-- Add uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
-- TODO: think about supporting all North America timezones
SET TIMEZONE="America/Indiana/Indianapolis";

-- Create Lessons Table
CREATE TABLE Lessons (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, -- lesson db ID
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (), -- lesson creation
    updated_at TIMESTAMP WITH TIME ZONE NULL, -- lesson update
    name VARCHAR (255) NOT NULL, -- lesson name
    lessonNumber VARCHAR (255) NOT NULL, -- lesson number, not sure if this is needed or not
    course VARCHAR (255) NOT NULL, -- course lesson belongs to
    active BOOLEAN NOT NULL, -- lesson still active or marked as "old"
    currentVersion TEXT NOT NULL, -- link to live lesson docs
    gradeRange int4range(0,12) NOT NULL, -- gradeRange lessons cover
    learningObjectives TEXT NOT NULL, -- lesson learning objectives
    sel BOOLEAN NOT NULL, -- is lesson SEL or nah?
    types NOT NULL,
    kStandards NULL,
    oneStandards NULL,
    twoStandards NULL,
    threeStandards NULL,
    fourStandards NULL,
    fiveStandards NULL,
    sixStandards NULL,
    sevenStandards NULL,
    eightStandards NULL,
    nineStandards NULL,
    tenStandards NULL,
    elevenStandards NULL,
    twelveStandards NULL,
)

CREATE INDEX active_lessons ON Lessons (name) WHERE active = TRUE;

CREATE TABLE Standards (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    standardId VARCHAR (255) NOT NULL,
    standard VARCHAR (255) NOT NULL,
    concept VARCHAR (255) NULL,
    subconcept VARCHAR (255) NULL,
)

CREATE INDEX active_standards ON Standards (standardId);

CREATE Courses ()

CREATE INDEX active_courses on Courses (courseId);

