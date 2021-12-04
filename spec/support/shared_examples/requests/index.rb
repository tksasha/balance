# frozen_string_literal: true

RSpec.shared_examples 'index.js' do
  it { should render_template :index }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end

RSpec.shared_examples 'index.html' do
  it { should render_template :index }

  it { expect(response.content_type).to eq 'text/html; charset=utf-8' }
end
