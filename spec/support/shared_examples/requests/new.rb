# frozen_string_literal: true

RSpec.shared_examples 'new.js' do
  it { should render_template :new }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end
