# frozen_string_literal: true

module Backoffice
  module Tags
    class GetResultService < ApplicationService
      def initialize(action_name, category, params)
        @action_name = action_name

        @category = category

        @params = params
      end

      def call
        return ::Backoffice::Tags::InitializeService.call(@category) if new?

        return ::Backoffice::Tags::GetResourceService.call(@params) if edit?

        return ::Backoffice::Tags::UpdateService.call(@params) if update?

        "::Backoffice::Tags::#{ @action_name.camelize }Service".constantize.call(@category, @params)
      end

      private

      def new?
        @action_name == 'new'
      end

      def edit?
        @action_name == 'edit'
      end

      def update?
        @action_name == 'update'
      end
    end
  end
end
