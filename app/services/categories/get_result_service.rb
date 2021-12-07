# frozen_string_literal: true

module Categories
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Categories::InitializeService.call if new?

      "::Categories::#{ @action_name.camelize }Service".constantize.call(@params)
    end

    private

    def new?
      @action_name == 'new'
    end
  end
end
