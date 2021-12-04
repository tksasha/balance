# frozen_string_literal: true

RSpec.shared_examples 'update.js' do
  it { should render_template :update }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end
