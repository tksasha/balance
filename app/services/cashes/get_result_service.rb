# frozen_string_literal: true

module Cashes
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Cashes::GetResourceService.call(@params) if show_or_edit?

      ::Cashes::UpdateService.call(@params)
    end

    private

    def show_or_edit?
      %w[show edit].include?(@action_name)
    end
  end
end
