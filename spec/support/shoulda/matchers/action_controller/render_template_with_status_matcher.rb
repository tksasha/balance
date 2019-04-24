module Shoulda
  module Matchers
    module ActionController
      class RenderTemplateMatcher
        def with_status status
          RenderTemplateWithStatusMatcher.new @options, status, @message, @context
        end
      end

      class RenderTemplateWithStatusMatcher
        def initialize template, status, message, context
          @render_template = RenderTemplateMatcher.new template, message, context

          @respond_with = RespondWithMatcher.new status
        end

        def matches? controller
          @controller = controller

          renders_template? && responds_with?
        end

        def description
          "#{ @render_template.description } and #{ @respond_with.description }"
        end

        def failure_message
          @render_template.failure_message || @respond_with.failure_message
        end

        def failure_message_when_negated
          "Did not expect to render #{ template } and respond with #{ status }"
        end

        private
        def renders_template?
          @render_template.matches? @controller
        end

        def responds_with?
          @respond_with.matches? @controller
        end

        def template
          @render_template.instance_variable_get :@template
        end

        def status
          @respond_with.instance_variable_get :@status
        end
      end
    end
  end
end
