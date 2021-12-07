# frozen_string_literal: true

module Categories
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      ::Categories::InitializeService.call
    end

    private

    def new?
      @action_name == 'new'
    end
  end
end
